package models

import (
	"context"
	"log"
	"time"
)

func (m *mysqlRepository) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs,
	            cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight,
		    lifespan, coalesce(details, ''),
		    coalesce(alternate_names, ''), coalesce(geographic_origin, '')
		    from dog_breeds order by breed`

	var breeds []*DogBreed

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b DogBreed
		err := rows.Scan(
			&b.ID,
			&b.Breed,
			&b.WeightLowLbs,
			&b.WeightHighLbs,
			&b.AverageWeightLbs,
			&b.LifeSpan,
			&b.Details,
			&b.AlternateNames,
			&b.GeographicOrigin,
		)
		if err != nil {
			return nil, err
		}
		breeds = append(breeds, &b)
	}

	return breeds, nil
}

func (m *mysqlRepository) GetBreedByName(b string) (*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs,
	            cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight,
		    lifespan, coalesce(details, ''),
		    coalesce(alternate_names, ''), coalesce(geographic_origin, '')
		    from dog_breeds where breed = ?`

	var breed DogBreed

	err := m.DB.QueryRowContext(ctx, query, b).Scan(
		&breed.ID,
		&breed.Breed,
		&breed.WeightLowLbs,
		&breed.WeightHighLbs,
		&breed.AverageWeightLbs,
		&breed.LifeSpan,
		&breed.Details,
		&breed.AlternateNames,
		&breed.GeographicOrigin,
	)
	if err != nil {
		return nil, err
	}

	return &breed, nil
}

func (m *mysqlRepository) GetDogOfMonthByID(id int) (*DogOfMonth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, video, image from dog_of_month where id = ?`

	var dom DogOfMonth

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&dom.ID,
		&dom.Video,
		&dom.Image,
	)
	if err != nil {
		log.Println("Error getting dog of the month by id:", err)
		return nil, err
	}

	return &dom, nil
}
