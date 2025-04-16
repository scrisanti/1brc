package main

import (
	// "1brc/logger"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/scrisanti/1brc/logger"
)

func main() {

	// CLI flags
	filename := flag.String("filename", "measurements_100M.txt", "Filename to the CSV file (in 'data' folder)")
	flag.Parse()

	// Validation
	if *filename == "" {
		fmt.Println("Error: --file is required")
		flag.Usage()
		os.Exit(1)
	}

	// Get User Home Dir
	usr, err := user.Current()
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't Get User: %v", err))
	}
	home := usr.HomeDir

	// Log To Standard Dir
	logDir := filepath.Join(home, "logs")
	logFilename := "log_1brc.log"
	logFP := filepath.Join(logDir, logFilename)
	// logFile, err := os.OpenFile(logFP, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	err = logger.InitLogger(logFP)
	if err != nil {
		panic(err)
	}
	// defer logFile.Close()

	// mw := io.MultiWriter(os.Stdout, logFile)
	//handler := slog.NewTextHandler(mw, &slog.HandlerOptions{Level: slog.LevelDebug})
	// logger := slog.New(handler)
	//slog.SetDefault(logger)

	slog.Info("# ----- 1brc ----- # ")
	cwd, err := os.Getwd()
	if err != nil {
		slog.Error(fmt.Sprintf("Error Getting CWD: %v", err))
	}
	slog.Debug("User CWD:", "cwd", cwd)

	// ---------------------- //
	start := time.Now()
	// output := bufio.NewWriter(logFile) // os.Stdout)
	dataFilepath := filepath.Join(home, "Documents", "data", *filename)
	st, err := os.Stat(dataFilepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	size := st.Size()

	err = baseline(dataFilepath)

	if err != nil {
		fmt.Println("Process Failed!")
	}

	// output.Flush()
	elapsed := time.Since(start)
	slog.Info("Processed File Successfully!", "size_MB",
		float64(size)/(1024*1024), "process_time_s", elapsed)

}
