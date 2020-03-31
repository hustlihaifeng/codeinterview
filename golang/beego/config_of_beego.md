# beego包
## beego/config.go文件

1. 定义两个关键全局变量：BConfig、 AppConfig

- BConfig是一个beego.Config对象，里面包含beego的一些默认配置变量，如 RunMode
- AppConfig是一个beegoAppConfig对象，里面的innerConfig是一个config.Configer这样一个接口。

```go
var (
	// BConfig is the default config for Application
	BConfig *Config
	// AppConfig is the instance of Config, store the config information from file
	AppConfig *beegoAppConfig

	// appConfigPath is the path to the config files
	appConfigPath string
	// appConfigProvider is the provider for the config, default is ini
	appConfigProvider = "ini"
)
```

2. init函数里面,先`BConfig = newBConfig()`来设置了默认变量，然后设置AppConfig。设置AppConfig时，先获取当前目录下面的conf/app.conf，没有则获取发起本命令下面的conf/app.conf(所以会导致go test的问题)。

- 如果这两个都不存在，则AppConfig.innerConfig被设置为config.fakeConfigContainer, config.fakeConfigContainer里面有`data map[string]string`，用来存储配置。
- 如果两个文件有任何一个存在，赋值给appConfigPath，然后`parseConfig(appConfigPath)`。感觉这里面包含对AppConfig和BConfig的设置代码。
- parseConfig里面首先调用`config.NewConfig(appConfigProvider, appConfigPath)`来得到一个config.Configer对象，然后将其赋值给beegoAppConfig.innerConfig. **这个`config.NewConfig`可以读取配置文件，然后返回一个Configer对象。详见[config包](#config包)**。 接着调用`assignConfig(AppConfig)`, assignConfig里面都是对一些beego的系统变量进行设置。前面的都是包内可见,是在beego的init函数里面就调用的逻辑，好像对我们没有多大用处。


```go
func init() {
	BConfig = newBConfig()
	var err error
	if AppPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath = filepath.Join(workPath, "conf", "app.conf")
	if !utils.FileExists(appConfigPath) {
		appConfigPath = filepath.Join(AppPath, "conf", "app.conf")
		if !utils.FileExists(appConfigPath) {
			AppConfig = &beegoAppConfig{innerConfig: config.NewFakeConfig()}
			return
		}
	}
	if err = parseConfig(appConfigPath); err != nil {
		panic(err)
	}
}
```

3. 包外可见的:LoadAppConfig

```go
// LoadAppConfig allow developer to apply a config file
func LoadAppConfig(adapterName, configPath string) error
````
里面更改了appConfigPath、appConfigProvider这些包变量，然后调用`parseConfig(appConfigPath)`来解析新的配置文件。parseConfig里面调用的assignConfig方法对BConfig的"RunMode"、"StaticDir"、"StaticExtensionsToGzip"、"LogOutputs"等部分进行了显示设置。parseConfig里面的`assignSingleConfig`对`BConfig, &BConfig.Listen, &BConfig.WebConfig, &BConfig.Log, &BConfig.WebConfig.Session`这些中的String、Int、Int64、Bool对象进行了设置。所以如果配置文件里面修改了，那么LoadAppConfig里面还是会对BConfig和AppConfig里面的所有的项目进行设置的。而且都是读取内存值，而不是重新读配置文件，那么我们从配置中心拉取配置是可行的。`LoadAppConfig`的**交互单位是文件**。如此说来，如果配置用文件存储的话，那么似乎还省事。但是如果放到db里面的话，后面似乎有更多可能。简单操作可以先发文件过去。接下来调研下有没有从web拉取配置来更新的方法。如果有一个Configer对象，似乎也可以先保存文件，后从文件读取来做。大致从beego/config.go和beego/config包，目前没有发现其他有效的方法可以从map或者web读取配置。不过config包里面有`IniConfigContainer.SaveConfigFile`这样一个样板函数，真要做的话，参照这个函数将map转换也是可以做的。

4. 包外可见的：SaveConfigFile调用了内部Configer的`SaveConfigFile`方法。
5. 包外可见的：有好多对beegoAppConfig的读取设置方法。

# config包
1. config.Config提供了读文件的接口，返回Configer对象。

```go
type Config interface {
	Parse(key string) (Configer, error)
	ParseData(data []byte) (Configer, error)
}
```

实现了Config接口的struct可以通过`Register`注册为一个新的adapter。

2. `NewConfig`可以从文件中读取，得到configer对象

```go
// NewConfig adapterName is ini/json/xml/yaml.
// filename is the config file path.
func NewConfig(adapterName, filename string) (Configer, error)
```

3. `NewConfigData`从字节数组中读取配置文件

```go
// NewConfigData adapterName is ini/json/xml/yaml.
// data is the config data.
func NewConfigData(adapterName string, data []byte) (Configer, error)
```
4. Configer里面有具体的配置存储方法

```go
// Configer defines how to get and set value from configuration raw data.
type Configer interface {
	Set(key, val string) error   //support section::key type in given key when using ini type.
	String(key string) string    //support section::key type in key string when using ini and json type; Int,Int64,Bool,Float,DIY are same.
	Strings(key string) []string //get string slice
	Int(key string) (int, error)
	Int64(key string) (int64, error)
	Bool(key string) (bool, error)
	Float(key string) (float64, error)
	DefaultString(key string, defaultVal string) string      // support section::key type in key string when using ini and json type; Int,Int64,Bool,Float,DIY are same.
	DefaultStrings(key string, defaultVal []string) []string //get string slice
	DefaultInt(key string, defaultVal int) int
	DefaultInt64(key string, defaultVal int64) int64
	DefaultBool(key string, defaultVal bool) bool
	DefaultFloat(key string, defaultVal float64) float64
	DIY(key string) (interface{}, error)
	GetSection(section string) (map[string]string, error)
	SaveConfigFile(filename string) error
}
```
对默认的几种adapter：ini/json/xml/yaml 来说，上面这些方法都已经实现。我们只需要给出adapter的名字，然后对应的文件名护着字节数组，就可以调用`NewConfig`或者`NewConfigData`来得到一个对应的Configer对象。这是一个公共库，剩下的就看beego如何使用这些库了。beego包的beegoAppConfig对象是对`config.Configer`的一个组合。