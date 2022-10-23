package domain_company_repository

import (
	"context"
	"net/http"

	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"

	. "github.com/felixa1996/go_next_be/app/infra/error"

	domain "github.com/felixa1996/go_next_be/app/domain/company"
)

func (r *companyMongoRepository) Upsert(ctx context.Context, company domain.Company) (domain.Company, error) {
	var companyFind domain.Company
	err := r.db.Database.Collection(domain.CollectionName).FindOne(context.TODO(), bson.M{"id": company.Id}).Decode(&companyFind)
	if err != nil {
		r.logger.Info("Company not found", zap.String("companyId", company.Id))
	}

	if len(companyFind.Id) > 0 {
		update := bson.M{
			"$set": bson.M{
				"company_name": company.CompanyName,
				"fieldbool":    false,
			},
		}
		r.logger.Info("Update company", zap.String("companyId", companyFind.Id))
		_, err = r.db.Database.Collection(domain.CollectionName).UpdateOne(context.TODO(), bson.M{"id": company.Id}, update)
		if err != nil {
			r.logger.Error("Failed to update company repository", zap.Error(err))
			return domain.Company{}, NewErrorWrapper(http.StatusInternalServerError, err, "Failed to update company repository")
		}
	}

	r.logger.Info("Insert new company", zap.String("companyName", companyFind.CompanyName))
	_, err = r.db.Database.Collection(domain.CollectionName).InsertOne(context.TODO(), company)
	if err != nil {
		r.logger.Error("Failed to insert company repository", zap.Error(err))
		return domain.Company{}, NewErrorWrapper(http.StatusInternalServerError, err, "Failed to insert company repository")
	}

	return company, nil
}
