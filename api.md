# 全局公共参数

**全局Header参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| 暂无参数 |

**全局Query参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| 暂无参数 |

**全局Body参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| 暂无参数 |

**全局认证方式**

> 无需认证

# 状态码说明

| 状态码 | 中文描述 |
| --- | ---- |
| 暂无参数 |

# student

> 创建人: 鳶天煢

> 更新人: 鳶天煢

> 创建时间: 2025-01-16 16:26:35

> 更新时间: 2025-01-16 16:30:35

```text
暂无描述
```

**目录Header参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| 暂无参数 |

**目录Query参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| 暂无参数 |

**目录Body参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| 暂无参数 |

**目录认证信息**

> 继承父级

## 保存学生信息

> 创建人: 鳶天煢

> 更新人: 鳶天煢

> 创建时间: 2025-01-16 16:26:35

> 更新时间: 2025-01-16 16:27:28

**保存学生的学号,姓名,性别,班级,成绩**

**接口状态**

> 开发中

**接口URL**

> /student/save

| 环境  | URL |
| --- | --- |


**请求方式**

> POST

**Content-Type**

> json

**请求Header参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Content-Type | application/json | String | 是 | - |

**请求Body参数**

```javascript
{
	"class": "",
	"gender": "",
	"name": "",
	"score": {},
	"student_id": ""
}
```

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| class | - | String | 否 | - |
| gender | - | String | 否 | - |
| name | - | String | 否 | - |
| score | - | Object | 否 | - |
| student_id | - | String | 否 | - |

**认证方式**

> 继承父级

**响应示例**

* OK(200)

```javascript
{
	"data": "",
	"message": "",
	"success": ""
}
```

| 参数名 | 示例值 | 参数类型 | 参数描述 |
| --- | --- | ---- | ---- |
| data | - | String | - |
| message | - | String | - |
| success | - | Boolean | - |

* 失败(404)

```javascript
暂无数据
```

* Bad Request(400)

```javascript
{
	"data": "",
	"message": "",
	"success": ""
}
```

| 参数名 | 示例值 | 参数类型 | 参数描述 |
| --- | --- | ---- | ---- |
| data | - | String | - |
| message | - | String | - |
| success | - | Boolean | - |

## 查询学生信息

> 创建人: 鳶天煢

> 更新人: 鳶天煢

> 创建时间: 2025-01-16 16:26:35

> 更新时间: 2025-01-16 16:27:30

**根据学生的学号查询学生信息**

**接口状态**

> 开发中

**接口URL**

> /student/get?id=

| 环境  | URL |
| --- | --- |


**请求方式**

> GET

**Content-Type**

> json

**请求Header参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Content-Type | application/json | String | 是 | - |

**请求Query参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| id | - | String | 是 | 学生ID |

**请求Body参数**

```javascript
暂无数据
```

**认证方式**

> 继承父级

**响应示例**

* OK(200)

```javascript
{
	"data": "",
	"message": "",
	"success": ""
}
```

| 参数名 | 示例值 | 参数类型 | 参数描述 |
| --- | --- | ---- | ---- |
| data | - | String | - |
| message | - | String | - |
| success | - | Boolean | - |

* Not Found(404)

```javascript
{
	"data": "",
	"message": "",
	"success": ""
}
```

| 参数名 | 示例值 | 参数类型 | 参数描述 |
| --- | --- | ---- | ---- |
| data | - | String | - |
| message | - | String | - |
| success | - | Boolean | - |

## 显示所有学生信息

> 创建人: 鳶天煢

> 更新人: 鳶天煢

> 创建时间: 2025-01-16 16:26:35

> 更新时间: 2025-01-16 16:26:35

**显示所有学生信息**

**接口状态**

> 开发中

**接口URL**

> /student/show

| 环境  | URL |
| --- | --- |


**请求方式**

> GET

**Content-Type**

> none

**请求Header参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Content-Type | application/json | String | 是 | - |

**认证方式**

> 继承父级

**响应示例**

* OK(200)

```javascript
{
	"data": "",
	"message": "",
	"success": ""
}
```

| 参数名 | 示例值 | 参数类型 | 参数描述 |
| --- | --- | ---- | ---- |
| data | - | String | - |
| message | - | String | - |
| success | - | Boolean | - |

* 失败(404)

```javascript
暂无数据
```

## 更新学生信息

> 创建人: 鳶天煢

> 更新人: 鳶天煢

> 创建时间: 2025-01-16 16:26:35

> 更新时间: 2025-01-16 16:26:35

**通过学生的 id 更新学生的详细信息**

```
id
```

**接口状态**

> 开发中

**接口URL**

> /student/update?id=

| 环境  | URL |
| --- | --- |


**请求方式**

> PUT

**Content-Type**

> json

**请求Header参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Content-Type | application/json | String | 是 | - |

**请求Query参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| id | - | String | 是 | 学生ID |

**请求Body参数**

```javascript
{
	"class": "",
	"gender": "",
	"name": "",
	"score": {},
	"student_id": ""
}
```

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| class | - | String | 否 | - |
| gender | - | String | 否 | - |
| name | - | String | 否 | - |
| score | - | Object | 否 | - |
| student_id | - | String | 否 | - |

**认证方式**

> 继承父级

**响应示例**

* OK(200)

```javascript
{
	"data": "",
	"message": "",
	"success": ""
}
```

| 参数名 | 示例值 | 参数类型 | 参数描述 |
| --- | --- | ---- | ---- |
| data | - | String | - |
| message | - | String | - |
| success | - | Boolean | - |

* Not Found(404)

```javascript
{
	"data": "",
	"message": "",
	"success": ""
}
```

| 参数名 | 示例值 | 参数类型 | 参数描述 |
| --- | --- | ---- | ---- |
| data | - | String | - |
| message | - | String | - |
| success | - | Boolean | - |

## 删除学生信息

> 创建人: 鳶天煢

> 更新人: 鳶天煢

> 创建时间: 2025-01-16 16:26:35

> 更新时间: 2025-01-16 16:27:18

**通过学生的 id 删除学生的详细信息**

```
id
```

**接口状态**

> 开发中

**接口URL**

> /student/delete?id=

| 环境  | URL |
| --- | --- |


**请求方式**

> DELETE

**Content-Type**

> json

**请求Header参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| Content-Type | application/json | String | 是 | - |

**请求Query参数**

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| --- | --- | ---- | ---- | ---- |
| id | - | String | 是 | 学生ID |

**请求Body参数**

```javascript
暂无数据
```

**认证方式**

> 继承父级

**响应示例**

* OK(200)

```javascript
{
	"data": "",
	"message": "",
	"success": ""
}
```

| 参数名 | 示例值 | 参数类型 | 参数描述 |
| --- | --- | ---- | ---- |
| data | - | String | - |
| message | - | String | - |
| success | - | Boolean | - |

* Not Found(404)

```javascript
{
	"data": "",
	"message": "",
	"success": ""
}
```

| 参数名 | 示例值 | 参数类型 | 参数描述 |
| --- | --- | ---- | ---- |
| data | - | String | - |
| message | - | String | - |
| success | - | Boolean | - |
