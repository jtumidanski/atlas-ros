package drop

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

func Initialize(l logrus.FieldLogger, db *gorm.DB) {
	d, e := os.LookupEnv("DROP_DIR")
	if !e {
		d = "/data/drops"
	}
	reactors, err := readDataDirectory(l, d)
	if err != nil {
		l.Fatal(err.Error())
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		err := deleteAll(tx)
		if err != nil {
			l.WithError(err).Errorf("Unable to truncate drops for initialization.")
		}

		for _, r := range reactors {
			for _, i := range r.Items {
				err := create(tx, r.Id, i.ItemId, i.QuestId, i.Chance)
				if err != nil {
					l.WithError(err).Errorf("Unable to insert drop %d for reactor %d.", i.ItemId, r.Id)
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		l.WithError(err).Errorf("Unable to initialize reactor drop database.")
	}
}

func GetByClassification(l logrus.FieldLogger, db *gorm.DB) func(classification uint32) []*Model {
	return func(classification uint32) []*Model {
		results, err := getByClassification(db)(classification)
		if err != nil {
			l.WithError(err).Warnf("Unable to retrieve drops for classification %d, proceeding with none.", classification)
			return make([]*Model, 0)
		}
		return results
	}
}
