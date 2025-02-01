package dto

type CreateItemInput struct {
	Name        string `json:"name" binding:"required,min=2"`
	Price       uint   `json:"price" binding:"required,min=1,max=999999"`
	Description string `json:"description"`
}

// 存在する値だけを更新する
// nil許容にするためポインタ型にする
type UpdateItemInput struct {
	Name        *string `json:"name" binding:"omitnil,min=2"` // nilではない場合は最小で2文字
	Price       *uint   `json:"price" binding:"omitnil,min=1,max=999999"`
	Description *string `json:"description"`
	SoldOut     *bool   `json:"soldOut"`
}
