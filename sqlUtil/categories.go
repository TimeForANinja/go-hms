package sqlUtil

import (
	"database/sql"
	"home_media_server/structs"
)

// GetCategorieByID returns a group from database by his id
func GetCategorieByID(db *sql.DB, id int) (*structs.Categorie, error) {
	rows, err := db.Query("SELECT * FROM `categories` WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []*structs.Categorie
	for rows.Next() {
		s := structs.Categorie{}
		err := rows.Scan(&s.ID, &s.Name, &s.ParentCategorie)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &s)
	}
	// sth went absolutely wrong when more than one with a UNIQUE id was found
	if len(categories) != 1 {
		return nil, nil
	}
	return categories[0], nil
}

// AddCategorie inserts a group into database, group.id will get set after the insert
func AddCategorie(db *sql.DB, categorie *structs.Categorie) (*structs.Categorie, error) {
	resp, err := db.Exec("INSERT INTO `groups` (name, parentCategorie) VALUES(?,?)", categorie.Name, categorie.ParentCategorie)
	if err != nil {
		return nil, err
	}
	lastInsert, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}
	categorie.ID = int(lastInsert)
	return categorie, nil
}
