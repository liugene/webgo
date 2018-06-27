package di

type Definition struct {
	/**
     * 实例别名
     * string alias
     */
	alias string
	/**
     * 是否单例
     * Bool isSingleton
     */
	isSingleton bool
	isEager bool
	callback func()
	className string
	isInstance bool
	instance struct{}
}

