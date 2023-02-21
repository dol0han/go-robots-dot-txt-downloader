Go Robots.txt Downloader

This is a simple Go program that downloads the /robots.txt file from a list of domains and saves it to a newly created folder with the same name as the domain.
Prerequisites

```sudo apt install golang-go```

To run this program, you need to have Go installed on your system. If you don't have Go installed, you can download it from the official Go website.
Usage

1. Edit the file named domains.txt in the same directory as the program.\n
2. Add the list of domains, one per line, to the domains.txt file.
3. Open a terminal and navigate to the directory where the program is located.
4. Run the command go run main.go to start the program.
5. The program will create a new directory for each domain in the domains.txt file and download the /robots.txt file to the directory.

Contributing

Contributions are welcome! If you find a bug or have an idea for a new feature, please open an issue or submit a pull request.
License

This program is released under the MIT License.
