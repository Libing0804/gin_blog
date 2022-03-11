package dao

import (
	"blog_gin/models"
	"errors"
	"log"
)
func GetCategoryById(cid int)(string,error){
	row:=DB.QueryRow("select name from blog_category where cid = ?",cid)
	var categoryName string
	if row.Err()!=nil{
		return categoryName,errors.New("查询分类名失败")
	}

	err:=row.Scan(&categoryName)
	if err!=nil{
		return categoryName,err
	}
	return categoryName,nil
}
func GetAllCategory()([]models.Category,error){
	rows,err:=DB.Query("select * from blog_category")
	if err!=nil{

		log.Println("category,error",err)
		return  nil,err
	}
	var categorys []models.Category
	for rows.Next(){
		var category models.Category
		err=rows.Scan(&category.Cid,&category.Name,&category.CreateAt,&category.UpdateAt)
		if err!=nil{
			return  nil,err
		}
		categorys=append(categorys,category)
	}
	return categorys,nil
}
