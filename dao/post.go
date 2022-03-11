package dao

import (
	"blog_gin/models"
	"log"
)

func GetAllPost() ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post ")
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func DeletePost(pid int) {
	_, err := DB.Exec("delete from blog_post where pid=? ", pid)
	if err != nil {
		log.Println(err)
		return
	}

}
func UpdatePost(post *models.Post) {
	_, err := DB.Exec("UPDATE blog_post set title=? ,content=?,markdown=?,category_id=?,type=?,slug=?, update_at=? where user_id=?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.UserId,
	)
	if err != nil {
		log.Println(err)

	}
}
func SavePost(post *models.Post) {
	ret, err := DB.Exec("insert into blog_post (title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) "+
		"values (?,?,?,?,?,?,?,?,?,?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt,
	)
	if err != nil {
		log.Println(err)

	}
	//获取自动更新的id
	pid, _ := ret.LastInsertId()
	post.Pid = int(pid)

}
func CountGetAllPostBySlug(slug string) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where slug=?", slug)
	_ = rows.Scan(&count)
	return
}
func CountGetAllPost() (count int) {
	rows := DB.QueryRow("select count(1) from blog_post")
	_ = rows.Scan(&count)
	return
}
func CountGetPostByCategory(cid int) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where category_id= ? ", cid)
	_ = rows.Scan(&count)
	return
}
func GetPostPageBySlug(slug string, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	row, err := DB.Query("select * from  blog_post where slug = ? limit ?,?", slug, page, pageSize)
	var posts []models.Post
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var post models.Post
		err = row.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func GetPostPageBySearch(search string) ([]models.Post, error) {

	row, err := DB.Query("select * from  blog_post where title  like ?", "%"+search+"%")
	var posts []models.Post
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var post models.Post
		err = row.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func GetPostById(pid int) (models.Post, error) {
	row := DB.QueryRow("select * from blog_post where pid = ? ", pid)
	var post models.Post
	if row.Err() != nil {
		log.Println("数据库查询文章失败")

		return post, row.Err()
	}

	err := row.Scan(&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt,
	)

	if err != nil {

		return post, err
	}
	return post, nil
}
func GetPostPage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	row, err := DB.Query("select * from  blog_post limit ?,?", page, pageSize)
	var posts []models.Post
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var post models.Post
		err = row.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil

}
func GetPostPageByCategory(cid, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	row, err := DB.Query("select * from  blog_post where category_id = ? limit ?,?", cid, page, pageSize)
	var posts []models.Post
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var post models.Post
		err = row.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
