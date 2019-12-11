const webpack = require("webpack")
module.exports = {
	publicPath: './',
	// devServer:{
	// 	proxy:{
	// 		'/api':{
	// 			target:'http://localhost:8081/Users/zt/Project',
	// 			changeOrigin:true,
	// 			pathRewrite:{
	// 				'^/api':''
	// 			}
	// 		}
	// 	}
	// },
	configureWebpack: {
		plugins: [
			new webpack.ProvidePlugin({
				$: "jquery",
				jQuery: "jquery",
				"windows.jQuery": "jquery"
			})
		]
	}
}
