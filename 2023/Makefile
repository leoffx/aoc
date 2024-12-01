
# The folder to create
FOLDER=day$(DAY)

# Default target
all:
	@echo "Specify a target."

# Target to create a new day's folder with the specified structure
new-day:
	@echo "Creating folder structure for day $(DAY)"
	mkdir $(FOLDER)
	touch $(FOLDER)/example.txt
	touch $(FOLDER)/input.txt
	cp template/main.go $(FOLDER)/main.go
	@echo "Folder structure for day $(DAY) created."



.PHONY: all new-day
