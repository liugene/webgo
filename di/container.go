package di

import (
	"reflect"
)

type Container struct {
	/**
     * 保存所有绑定的实例
     * key 为alias value 为 InstanceDefinition实例对象
     * @Array $_service
     */
	definitionMap []Definition
}

func (c *Container) Bind(definition Definition) {
	/*
	   断言别名是否合法
	 */
	assertAliasNameAvailable(definition.alias)
	/*
	   存放注入
	 */
	c.definitionMap = []Definition{definition.alias:definition}
	/*
	   是否立即执行
	 */
	initEagerDefinition(definition)
}

func initEagerDefinition(definition Definition) struct{} {

	/*
	 * 如果是立即初始化即立即实例化
	 */
	if definition.isEager {
		/*
		 * 查看是否已经实例化过
		 */
		if definition.isInstance {
			return definition.instance
		}
	}

	return Container{}.Get(definition.alias)

}

func (c *Container) Get(alias string) struct{} {

	return Definition{}.instance

}

func (c *Container) Make(alias string) {

}

/*
 * 断言别名是否合法
 */
func assertAliasNameAvailable(alias string)  {

}
