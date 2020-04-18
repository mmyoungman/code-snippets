const path = require("path");

module.exports = {
  entry: "./src/index.tsx",
  output: {
    filename: "./dist/bundle.js"
  },
  devtool: "source-map",
  resolve: { extensions: ["*", ".ts", ".tsx", ".js"] },
  module: {
    rules: [
      { 
        test: /\.tsx?$/, 
        exclude: /node_modules/,
        use: "ts-loader" 
      }
    ]
  },
  devServer: {
    contentBase: path.join(__dirname, "public/"),
    port: 8080
  }
};
