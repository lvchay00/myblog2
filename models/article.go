package models

import (
	"fmt"
	_ "reflect"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id         int        `orm:"pk;auto"`                     //主键
	Title      string     `orm:"unique"`                      //索引独一无二
	Content    string     `orm:"type(text);null"`             //文本
	InTime     time.Time  `orm:"auto_now_add;type(datetime)"` //日期
	User       *User      `orm:"rel(fk)"`                     //外键
	Section    []*Section `orm:"rel(m2m)"`                    //外键 多对多
	View       int        `orm:"default(0)"`
	ReplyCount int        `orm:"default(0)"`
	Type       int        `orm:"default(0)"` //文章类型
}

func SaveArticle(Article *Article) (error, int64) {
	o := orm.NewOrm()
	id, err := o.Insert(Article)
	fmt.Println("SAVE Article", err)
	return err, id
}

func FindArticleById(id int) Article {
	o := orm.NewOrm()
	var Article Article
	o.QueryTable(Article).RelatedSel().Filter("Id", id).One(&Article)
	return Article
}
func FindArticleById_P(id int) *Article {
	o := orm.NewOrm()
	var Article Article
	o.QueryTable(Article).RelatedSel().Filter("Id", id).One(&Article)
	return &Article
}

func IncrView(Article *Article) {
	o := orm.NewOrm()
	Article.View = Article.View + 1
	o.Update(Article, "View")
}

func IncrReplyCount(Article *Article) {
	o := orm.NewOrm()
	Article.ReplyCount = Article.ReplyCount + 1
	o.Update(Article, "ReplyCount")
}

func ReduceReplyCount(Article *Article) {
	o := orm.NewOrm()
	Article.ReplyCount = Article.ReplyCount - 1
	o.Update(Article, "ReplyCount")
}

func FindArticleByUser(user *User, limit int) []*Article {
	o := orm.NewOrm()
	var article Article
	var articles []*Article
	o.QueryTable(article).RelatedSel().Filter("User", user).OrderBy("-LastReplyTime", "-InTime").Limit(limit).All(&articles)
	return articles
}

func UpdateArticle(Article *Article) {
	o := orm.NewOrm()
	o.Update(Article)
}

func DeleteArticle(Article *Article) {
	o := orm.NewOrm()
	o.Delete(Article)
}

func DeleteArticleByUser(user *User) {
	o := orm.NewOrm()
	o.Raw("delete from Article where user_id = ?", user.Id).Exec()
}

//SELECT * FROM myblog.article_sections
func DeleteArticleSectionsByArticleId(article_id int64) {
	o := orm.NewOrm()
	o.Raw("delete from article_sections where article_id = ?", article_id).Exec()
}

func SaveArticleSections(article_id int64, section_id int64) {
	o := orm.NewOrm()
	o.Raw("insert into article_sections (article_id, section_id) values (?, ?)", article_id, section_id).Exec()
}
func FindArticleById_2(id int) Article {
	o := orm.NewOrm()
	var Article Article
	o.QueryTable(Article).RelatedSel().Filter("Id", id).One(&Article)
	return Article
}

//根据类型找出所有的文章
func FindArticleByArticleSectionsUsingSection(p int, size int, section_id int) (int64, []Article) {
	// fmt.Println("测试开始")
	// FindArticleUsingSection(p, size, section_id)
	// fmt.Println("测试结束")
	o := orm.NewOrm()
	var article Article
	var articles []Article
	var res []orm.Params
	var total int64 = 0
	if section_id == 0 {

		qs := o.QueryTable(article)
		// if section.Id > 0 {
		// 	qs = qs.Filter("Section", section)
		// }
		total, _ = qs.Limit(-1).Count()
		qs.RelatedSel().OrderBy("-InTime").Limit(size).Offset((p - 1) * size).All(&articles)

	} else {
		//	o.Raw("select * from article_sections where  section_id = ?  limit ? OFFSET ?;", section_id, size, (p-1)*size).Values(&res, "id", "article_id", "section_id")
		//	o.Raw("select article_id from article_sections where  section_id = ?  limit ? OFFSET ?;", section_id, size, (p-1)*size).Values(&res, "id", "article_id", "section_id")

		o.Raw(`select article_id from article left join article_sections 
	         on article_sections.section_id = ? and article_sections.article_id = article.id
	         where article_sections.section_id  is not null order by article.in_time desc limit ? OFFSET ?;`,
			section_id, size, (p-1)*size).Values(&res, "article_id")
		fmt.Println("res", res)

		for _, v := range res {
			fmt.Println("v[\"article_id\"]", v["article_id"])

			value, ok := v["article_id"].(string)
			if ok {
				fmt.Println("v[\"article_id\"]的值", value)
				//转换为int 类型
				int_str, _ := strconv.Atoi(value)
				if int_str != 0 {
					article = (FindArticleById(int_str))
				//	article.Section = FindSectionByArticle(int64(int_str))
					articles = append(articles, article)
					total = total + 1
				}
			}
		}
	}
	return total, articles
}

//根据类型找出所有的文章
func FindArticleUsingSection(p int, size int, section_id int64) {
	o := orm.NewOrm()

	var res []orm.Params

	o.Raw(`select Article.* from Article left join article_sections 
	on article_sections.section_id = ? and article_sections.article_id = Article.id
	where article_sections.section_id  is not null order by Article.in_time desc limit ? OFFSET ?;`, section_id, size, (p-1)*size).
		Values(&res, "id", "title", "content", "in_time", "user_id", "view", "reply_count", "last_reply_user_id", "last_reply_time", "type")
	fmt.Println("res", res)

}
