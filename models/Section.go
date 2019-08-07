package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type Section struct {
	Id      int        `orm:"pk;auto"`
	Name    string     `orm:"unique"`
	Article []*Article `orm:"reverse(many)"`
}

func FindAllSection() []*Section {
	o := orm.NewOrm()
	var section Section
	var sections []*Section
	o.QueryTable(section).OrderBy("id").All(&sections)
	return sections
}
func FindSectionById(id int64) Section {
	o := orm.NewOrm()
	var section Section
	o.QueryTable(section).Filter("id", id).One(&section)
	return section
}
func FindSectionByArticle(article_id int64) []*Section {
	o := orm.NewOrm()

	var sections []*Section

	var res []orm.Params
	o.Raw("select section_id from article_sections where  article_id = ? ;", article_id).Values(&res, "section_id")
	fmt.Println("res", res)

	for _, v := range res {
		value, ok := v["section_id"].(string)
		if ok {
			fmt.Println("v[\"section_id\"]的值", value)
			//转换为int 类型
			int_str, _ := strconv.Atoi(value)
			if int_str != 0 {
				var s = (FindSectionById(int64(int_str)))
				sections = append(sections, &s)
			}
		}
	}
	return sections
}
func DeleteArticleSectionByArticle(article_id int64) {
	o := orm.NewOrm()
	var res []orm.Params
	o.Raw("select section_id from article_sections where  article_id = ? ;", article_id).Values(&res, "section_id")
	fmt.Println("res", res)

	for _, v := range res {
		value, ok := v["section_id"].(string)
		if ok {
			fmt.Println("v[\"section_id\"]的值", value)
			//转换为int 类型
			int_str, _ := strconv.Atoi(value)
			if int_str != 0 {
				o.Raw("delete from article_sections where  article_id = ? ;", article_id).Exec()
			}
		}
	}
}
func InsertSection(section *Section) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(section)
	return id
}
func DeleteSection(section *Section) {
	o := orm.NewOrm()
	o.Delete(section)
}
