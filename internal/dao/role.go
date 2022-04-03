package dao

import (
	"shippo-server/internal/model"
)

type RoleDao struct {
	*Dao
}

func NewRoleDao(s *Dao) *RoleDao {
	return &RoleDao{s}
}

// 新建角色,用name和remark创建
func (t *RoleDao) RoleCreate(name string, Remark string) (r model.Role, err error) {
	r.RoleName = name
	r.Remark = Remark
	err = t.db.Create(&r).Error
	return
}

// 删除角色
func (t *RoleDao) RoleDelete(id uint) (err error) {
	err = t.db.Where("id", id).Delete(&model.Role{}).Error
	return
}

// 修改角色名和备注
func (t *RoleDao) RoleUpdate(param model.Role) (err error) {
	err = t.db.Model(&model.Role{}).Select("role_name", "remark", "updated_at").Updates(&param).Error
	return
}

// 根据角色id查询所拥有的全部权限策略关联信息
func (t *RoleDao) RoleAssociationFind(roleId uint) (list []model.RoleAssociation, err error) {
	err = t.db.Where("role_id", roleId).Find(&list).Error
	return
}

// 根据角色id查询所拥有的全部权限策略关联信息,仅返回policy_id
func (t *RoleDao) RoleAssociationFindPolicyIdList(roleId uint) (list []uint, err error) {
	err = t.db.Model(&model.RoleAssociation{}).Select("policy_id").
		Where("role_id", roleId).Find(&list).Error
	return
}

// 创建角色关联
func (t *RoleDao) RoleAssociationCreate(roleId uint, policyId uint) (r model.RoleAssociation, err error) {
	r.RoleId = roleId
	r.PolicyId = policyId
	err = t.db.Create(&r).Error
	return
}

// 根据id删除角色关联
func (t *RoleDao) RoleAssociationDelById(id uint) (err error) {
	err = t.db.Where("id", id).Delete(&model.RoleAssociation{}).Error
	return
}

// 根据关联信息删除角色关联
func (t *RoleDao) RoleAssociationDel(roleId uint, policyId uint) (err error) {
	err = t.db.Where("role_id", roleId).Where("policy_id", policyId).
		Delete(&model.RoleAssociation{}).Error
	return
}

// 查询全部角色
func (t *RoleDao) RoleFindAll() (r []model.Role, err error) {
	err = t.db.Find(&r).Error
	return
}

// 查询某个角色所拥有的权限策略
func (t *RoleDao) RoleFindPermissionPolicy(id uint) (p []model.PermissionPolicy, err error) {
	subQuery := t.db.Model(&model.RoleAssociation{}).Select("policy_id").Where("role_id", id)
	err = t.db.Where("id IN (?)", subQuery).Find(&p).Error
	return
}

// 查询某个角色所拥有的访问规则
func (t *RoleDao) RoleFindPermissionAccess(id uint) (p []model.PermissionAccess, err error) {
	subQuery1 := t.db.Model(&model.RoleAssociation{}).Select("policy_id").Where("role_id", id)
	subQuery2 := t.db.Model(&model.PermissionAssociation{}).Select("access_id").
		Where("policy_id IN (?)", subQuery1)
	err = t.db.Where("id IN (?)", subQuery2).Find(&p).Error
	return
}

// 按照类型查询某个角色所拥有的访问规则
func (t *RoleDao) RoleFindPermissionAccessByType(id uint, accessType string) (
	p []model.PermissionAccess, err error) {
	subQuery1 := t.db.Model(&model.RoleAssociation{}).Select("policy_id").Where("role_id", id)
	subQuery2 := t.db.Model(&model.PermissionAssociation{}).Select("access_id").
		Where("policy_id IN (?)", subQuery1)
	err = t.db.Where("id IN (?)", subQuery2).Where("access_type", accessType).Find(&p).Error
	return
}
