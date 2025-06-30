package controllers

//维度	    POST vs PUT vs DELETE
//传参方式   三者完全一致，Gin 的绑定方法通用
//语义差异	POST 创建，PUT 替换，DELETE 删除
//幂等性	    PUT/DELETE 是幂等的，POST 不是
//请求体使用	DELETE 通常避免请求体，PUT 需要完整数据
//路由设计	DELETE 优先用路径参数，PUT/POST 灵活使用 JSON/表单
//最终结论：
//在 Gin 的技术实现上，DELETE 和 PUT 的传参与 POST 完全一致，但根据 RESTful 规范，
//它们的使用场景和设计意图不同，需要遵循各自的语义约束。
