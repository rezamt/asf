package swagger

//go:generate rm -rf
//go:generate mkdir -p efood
//go:generate swagger generate server --quiet --target efood --name efood --spec swagger.yml --exclude-main
