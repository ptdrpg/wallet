package repository

import "github.com/ptdrpg/wallet/entity"

func (r *Repository) FindAllTransaction() ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	if err := r.DB.Model(&entity.Transaction{}).Order("created_at").Find(&transactions).Error; err != nil {
		return []entity.Transaction{}, err
	}
	return transactions, nil
}

func (r *Repository) FindTransactionByUID(uid string) (entity.Transaction, error) {
	var transaction entity.Transaction
	if err := r.DB.Where("uid = ?", uid).Find(&transaction).Error; err != nil {
		return entity.Transaction{}, err
	}
	return transaction, nil
}

func (r *Repository) CreateTransaction(transaction *entity.Transaction) error {
	if err := r.DB.Create(transaction).Error; err != nil {
		return err
	}

	return nil
}
