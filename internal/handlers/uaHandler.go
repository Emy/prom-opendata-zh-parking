package handlers

import "math/rand"

// It's a Surprise Tool That Will Help Us Later
// ,n888888n,
// .8888888888b
// 888888888888nd8P~”8g,
// 88888888888888   _  `'~\.  .n.
// `Y888888888888. / _  |~\\ (8"8b
// ,nnn.. 8888888b.  |  \ \m\|8888P
// ,d8888888888888888b. \8b|.\P~ ~P8~
// 888888888888888P~~_~  `8B_|      |
// ~888888888~'8'   d8.    ~      _/
// ~Y8888P'   ~\ | |~|~b,__ __--~
// --~~\   ,d8888888b.\`\_/ __/~
// \_ d88888888888b\_-~8888888bn.
//
//	\8888P   "Y888888888888"888888bn.
//
// /~'\_"__)      "d88888888P,-~~-~888
// /  / )   ~\     ,888888/~' /  / / 8'
// (  / / / |) )   /   '"88(/ ~  / /  |
// (       /_/  /~        \( _/      /
// (_(_ ( /~~\/      ,  O,/~\___/_/'
//
//	  ~~~    |       \_  (
//			 )(        \_|
//		__--~"mb  ,g8888b.
//	  _/    8888b(.8P"~'~---__
//	 /       ~~~| / ,/~~~~--, `\
//
//	(       ~\,_) (/         ~-_`\
//	 \  -__---~._ \             ~\\
//	 (           )\\              ))
//	 `\          )  "-_           `|
//	   \__    __/      ~-__   __--~
//		  ~~"~             ~~~
func randomUA() string {
	var userAgents = []string{
		"Mozilla/5.0 (Linux; Android 10; SM-N976V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Mobile Safari/537.36",
		"Dalvik/1.6.0 (Linux; U; Android 4.4.4; Smart TV Build/20190812_185713)",
		"RubyBrowser/45.11.1 (iPhone; iOS 14.2; Scale/3.00)",
		"Mozilla/5.0 (Linux; Android 9; ASUS_X00TD) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.141 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 8.1.0; Primo G8i 4G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Mobile Safari/537.36",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; .NET CLR 1.1.4322; .NET CLR 3.0.04506.30; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET4.0C; .NET4.0E; InfoPath.2)",
		"Mozilla/5.0 (Linux; Android 10; LM-K200 Build/QKQ1.200311.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/81.0.4044.138 Mobile Safari/537.36",
		"Opera/9.80 (Android; Opera Mini/51.0.2254/187.20; U; en) Presto/2.12.423 Version/12.16",
		"Mozilla/5.0 (Linux; Android 9; SM-N9750) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Mobile Safari/537.36 OPR/58.2.2878.53403",
		"Dalvik/2.1.0 (Linux; U; Android 6.0; Mione XS max Build/MRA58K)",
	}

	return userAgents[rand.Intn(len(userAgents))]
}
