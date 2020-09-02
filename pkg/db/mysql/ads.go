package mysql

import (
	"log"

	"github.com/tle-dieu/gql_test/graph/models"
)

func (cli *ClientMySQL) SaveAd(ad models.AdInput) error {
	stmt, err := cli.db.Prepare("INSERT INTO Ads(ref,brand,model,price,bluetooth,gps) VALUES(?,?,?,?,?,?)")
	// should not fatal if duplicate ref
	if err != nil {
		return err
	}
	_, err = stmt.Exec(ad.Ref, ad.Brand, ad.Model, ad.Price, ad.Options.Bluetooth, ad.Options.Gps)
	if err != nil {
		return err
	}
	log.Println("Row inserted!")
	return nil
}

func (cli *ClientMySQL) UpdateAd(ad models.AdInput) error {
	stmt, err := cli.db.Prepare("UPDATE Ads SET brand = ?, model = ?, price = ?, bluetooth = ?, gps = ? WHERE ref = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(ad.Brand, ad.Model, ad.Price, ad.Options.Bluetooth, ad.Options.Gps, ad.Ref)
	if err != nil {
		return err
	}
	log.Println("Row updated!")
	return nil
}

func (cli *ClientMySQL) DeleteAd(ref string) error {
	stmt, err := cli.db.Prepare("DELETE FROM Ads WHERE ref = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(ref)
	if err != nil {
		return err
	}
	log.Println("Row deleted!")
	return nil
}

func (cli *ClientMySQL) GetAllAds() ([]models.Ad, error) {
	stmt, err := cli.db.Prepare("SELECT ref,brand,model,price,bluetooth,gps FROM Ads")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ads []models.Ad
	for rows.Next() {
		var ad models.Ad
		ad.Options = &models.Options{
			Bluetooth: new(bool),
			Gps:       new(bool),
		}
		err := rows.Scan(&ad.Ref, &ad.Brand, &ad.Model, &ad.Price, &ad.Options.Bluetooth, &ad.Options.Gps)
		if err != nil {
			return nil, err
		}
		ads = append(ads, ad)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ads, nil
}
