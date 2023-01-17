# VueJS



## Vue3 安装

### 1、独立版本

我们可以在 Vue.js 的官网上直接下载最新版本, 并用` <script> `标签引入。

[下载Vue.js](https://unpkg.com/vue@3.2.36/dist/vue.global.js)



### 2、使用 CDN 方法

以下推荐国外比较稳定的两个 CDN，国内还没发现哪一家比较好，目前还是建议下载到本地。

- **Staticfile CDN（国内）** : https://cdn.staticfile.org/vue/3.0.5/vue.global.js
- **unpkg**：https://unpkg.com/vue@next, 会保持和 npm 发布的最新的版本一致。
- **cdnjs** : https://cdnjs.cloudflare.com/ajax/libs/vue/3.0.5/vue.global.js



------------------------------------



## Vue3 创建项目



### vue create 命令

vue create 命令创建项目语法格式如下：

```vue
vue create [options] <app-name>
```

options 选项可以是：

- **-p, --preset <presetName>**： 忽略提示符并使用已保存的或远程的预设选项
- **-d, --default**： 忽略提示符并使用默认预设选项
- **-i, --inlinePreset <json>**： 忽略提示符并使用内联的 JSON 字符串预设选项
- **-m, --packageManager <command>**： 在安装依赖时使用指定的 npm 客户端
- **-r, --registry <url>**： 在安装依赖时使用指定的 npm registry
- **-g, --git [message]**： 强制 / 跳过 git 初始化，并可选的指定初始化提交信息
- **-n, --no-git**： 跳过 git 初始化
- **-f, --force**： 覆写目标目录可能存在的配置
- **-c, --clone**： 使用 git clone 获取远程预设选项
- **-x, --proxy**： 使用指定的代理创建项目
- **-b, --bare**： 创建项目时省略默认组件中的新手指导信息
- **-h, --help**： 输出使用帮助信息

接下来我们创建 runoob-vue3-app 项目：

```vue
vue create runoob-vue3-app
```

执行以上命令会出现安装选项界面：

```vue
Vue CLI v4.4.6
? Please pick a preset: (Use arrow keys)
❯ default (babel, eslint)
  Manually select features
```

按下回车键后就会进入安装，等候片刻即可完成安装。

安装完成后，我们进入项目目录：

cd runoob-vue3-app

启动应用：

```vue
npm run serve
```



### vue ui 命令

除了使用 **vue create** 命令创建项目，我们还可以使用可视化创建工具来创建项目。

```vue
vue ui
```

执行以上命令，会在浏览器弹出一个项目管理的界面：

<img src="https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301132127654.jpeg" alt="img" style="zoom: 33%;" /><img src="https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301132127961.jpeg" alt="img" style="zoom: 33%;" />

我们可以点击**"创建"**选项来创建一个项目，选择底部"在此创建项目"，页面上方也可以选择路径：

然后输入我们的项目名称，选择包管理工具为 npm，然后点击下一步：

<img src="https://raw.githubusercontent.com/disturb-yy/study-coding/main/img/202301132129921.jpeg" alt="img" style="zoom:33%;" /><img src="https://www.runoob.com/wp-content/uploads/2021/12/69A83D7A-C7FB-478B-8DA0-40BF673F160F.jpeg" alt="img" style="zoom:25%;" />

配置选择默认即可:

接下来就等待完成安装，安装完成管理界面如下：

<img src="https://www.runoob.com/wp-content/uploads/2021/12/4AE552A2-2AE4-4B23-AECA-90CE7D29C047-scaled.jpeg" alt="img" style="zoom:33%;" />



--------------------------



### Vue3 目录结构

上一章节中我们使用了 npm 安装项目(Vue-cli 和 Vite)，我们在 IDE（Vscode、Atom等） 中打开该目录，结构如下所示：



#### 目录解析

| 目录/文件    | 说明                                                         |
| :----------- | :----------------------------------------------------------- |
| build        | 项目构建(webpack)相关代码                                    |
| config       | 配置目录，包括端口号等。我们初学可以使用默认的。             |
| node_modules | npm 加载的项目依赖模块                                       |
| src          | 这里是我们要开发的目录，基本上要做的事情都在这个目录里。里面包含了几个目录及文件：assets: 放置一些图片，如logo等。components: 目录里面放了一个组件文件，可以不用。App.vue: 项目入口文件，我们也可以直接将组件写这里，而不使用 components 目录。main.js: 项目的核心文件。index.css: 样式文件。 |
| static       | 静态资源目录，如图片、字体等。                               |
| public       | 公共资源目录。                                               |
| test         | 初始测试目录，可删除                                         |
| .xxxx文件    | 这些是一些配置文件，包括语法配置，git配置等。                |
| index.html   | 首页入口文件，你可以添加一些 meta 信息或统计代码啥的。       |
| package.json | 项目配置文件。                                               |
| README.md    | 项目的说明文档，markdown 格式                                |
| dist         | 使用 **npm run build** 命令打包后会生成该目录。              |





## Vue3 模板语法



### 插值



#### 文本

数据绑定最常见的形式就是使用 **{{...}}**（双大括号）的文本插值：

```vue
<div id="app">
  <p>{{ message }}</p>
</div>
```

**{{...}}** 标签的内容将会被替代为对应组件实例中 **message** 属性的值，如果 **message** 属性的值发生了改变，**{{...}}** 标签内容也会更新。

如果不想改变标签的内容，可以通过使用 **v-once** 指令执行一次性地插值，当数据改变时，插值处的内容不会更新。

```vue
<span v-once>这个将不会改变: {{ message }}</span>
```



#### Html

使用 v-html 指令用于输出 html 代码：

```vue
<div id="example1" class="demo">
    <p>使用双大括号的文本插值: {{ rawHtml }}</p>
    <p>使用 v-html 指令: <span v-html="rawHtml"></span></p>
</div>
 
<script>
const RenderHtmlApp = {
  data() {
    return {
      rawHtml: '<span style="color: red">这里会显示红色！</span>'
    }
  }
}
 
Vue.createApp(RenderHtmlApp).mount('#example1')
</script>
```



#### 属性

HTML 属性中的值应使用 v-bind 指令。

```vue
<div v-bind:id="dynamicId"></div>
```



