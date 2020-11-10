package logger

// Init use logrus to logger.
// here can custom logrus.
// then can use logger.Info() in this project.
func Init() {
	// Log as JSON instead of the default ASCII formatter.
	//logger.SetFormatter(&logger.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//logger.SetOutput(os.Stdout)

	// Only logger the warning severity or above.
	//logger.SetLevel(logger.WarnLevel)
}
