# GolangStudy
项目初始化
    环境安装

        go get github.com/gin-gonic/gin
        go get -u gorm.io/gorm
        go get -u gorm.io/driver/mysql
        go get github.com/golang-jwt/jwt/v5
        go get github.com/spf13/viper
        go get github.com/gin-gonic/gin/binding

数据库设计与模型定义

    users 表：存储用户信息，包括 id 、 username 、 password 、 email 等字段。
    posts 表：存储博客文章信息，包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、 created_at 、 updated_at 等字段。
    comments 表：存储文章评论信息，包括 id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段。
    使用 GORM 定义对应的 Go 模型结构体。

用户认证与授权

    实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码。
    使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。

文章管理功能

    实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
    实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
    实现文章的更新功能，只有文章的作者才能更新自己的文章。
    实现文章的删除功能，只有文章的作者才能删除自己的文章。

评论功能

    实现评论的创建功能，已认证的用户可以对文章发表评论。
    实现评论的读取功能，支持获取某篇文章的所有评论列表。

错误处理与日志记录

    对可能出现的错误进行统一处理，如数据库连接错误、用户认证失败、文章或评论不存在等，返回合适的 HTTP 状态码和错误信息。
    使用日志库记录系统的运行信息和错误信息，方便后续的调试和维护。

