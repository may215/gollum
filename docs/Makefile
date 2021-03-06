# -*- makefile -*-

#
# Configuration
#

# General stuff
RST_GENERATOR_DIR=./generator
RST_GENERATOR=$(RST_GENERATOR_DIR)/rst_generator
PRINTF=/usr/bin/printf

# Locations of module sources
MODULES_DIR_PRODUCER=../producer
MODULES_DIR_CONSUMER=../consumer
MODULES_DIR_FILTER=../filter
MODULES_DIR_FORMATTER=../format
MODULES_DIR_ROUTER=../router

# Target directories for RST
TARGET_ROOT=src/gen
TARGET_DIR_PRODUCER=$(TARGET_ROOT)/producer
TARGET_DIR_CONSUMER=$(TARGET_ROOT)/consumer
TARGET_DIR_FILTER=$(TARGET_ROOT)/filter
TARGET_DIR_FORMATTER=$(TARGET_ROOT)/formatter
TARGET_DIR_ROUTER=$(TARGET_ROOT)/router

# Excluded files
EXCLUDE_NO_OP='^$$'
EXCLUDE_PRODUCER='(InfluxDBWriter[0-9][0-9]|filestate|kafkaMurmur2HashPartitioner|s3buffer|spoolfile)'
EXCLUDE_CONSUMER='proxyclient'
EXCLUDE_FILTER='regexp'
EXCLUDE_FORMATTER='collectddata'
EXCLUDE_ROUTER=$(EXCLUDE_NO_OP)

# Create lists of .rst files based on .go sources
rsts_producer  := $(shell find $(MODULES_DIR_PRODUCER)  -type f -name '*.go' -not -name '*_test.go' | grep -vE $(EXCLUDE_PRODUCER)  | sed s^$(MODULES_DIR_PRODUCER)^$(TARGET_DIR_PRODUCER)^   | sed 's|\.go$$|.rst|' | sort)
rsts_consumer  := $(shell find $(MODULES_DIR_CONSUMER)  -type f -name '*.go' -not -name '*_test.go' | grep -vE $(EXCLUDE_CONSUMER)  | sed s^$(MODULES_DIR_CONSUMER)^$(TARGET_DIR_CONSUMER)^   | sed 's|\.go$$|.rst|' | sort)
rsts_filter    := $(shell find $(MODULES_DIR_FILTER)    -type f -name '*.go' -not -name '*_test.go' | grep -vE $(EXCLUDE_FILTER)    | sed s^$(MODULES_DIR_FILTER)^$(TARGET_DIR_FILTER)^       | sed 's|\.go$$|.rst|' | sort)
rsts_formatter := $(shell find $(MODULES_DIR_FORMATTER) -type f -name '*.go' -not -name '*_test.go' | grep -vE $(EXCLUDE_FORMATTER) | sed s^$(MODULES_DIR_FORMATTER)^$(TARGET_DIR_FORMATTER)^ | sed 's|\.go$$|.rst|' | sort)
rsts_router    := $(shell find $(MODULES_DIR_ROUTER)    -type f -name '*.go' -not -name '*_test.go' | grep -vE $(EXCLUDE_ROUTER)    | sed s^$(MODULES_DIR_ROUTER)^$(TARGET_DIR_ROUTER)^       | sed 's|\.go$$|.rst|' | sort)

####################################################################################################

#
# Targets
#
all_groups := producers consumers filters formatters routers
all_preps  := prep_producers prep_consumers prep_filters prep_formatters prep_routers

.PHONY: all plugin_rst $(all_groups) $(all_preps) html

all: clean groups

all-really: clean clean-html all html

groups: $(all_groups)
preps:  $(all_preps)

clean:
	if [ -e $(RST_GENERATOR) ]; then rm $(RST_GENERATOR) ; fi
	find $(TARGET_ROOT) -type f -name '*.rst' -print -delete

$(BUILD_DIR):
	mkdir $@

$(RST_GENERATOR): $(BUILD_DIR)
	cd $(RST_GENERATOR_DIR) ; go build -o ../$(RST_GENERATOR)

# Subtargets
producers:  prep_producers  $(rsts_producer)
consumers:  prep_consumers  $(rsts_consumer)
filters:    prep_filters    $(rsts_filter)
formatters: prep_formatters $(rsts_formatter)
routers:    prep_routers    $(rsts_router)

# Prepare target dirs
prep_producers:
	mkdir -p $(TARGET_DIR_PRODUCER)
	if [ -e $(TARGET_DIR_PRODUCER)/index.rst ]; then rm $(TARGET_DIR_PRODUCER)/index.rst ; fi

	$(PRINTF) ".. Autogenerated by docs/Makefile\n"         >> $(TARGET_DIR_PRODUCER)/index.rst
	$(PRINTF) ".. include:: ../../plugins/producer.rst\n" >> $(TARGET_DIR_PRODUCER)/index.rst
	$(PRINTF) "\n"                 >> $(TARGET_DIR_PRODUCER)/index.rst
	$(PRINTF) ".. toctree::\n"     >> $(TARGET_DIR_PRODUCER)/index.rst
	$(PRINTF) "    :maxdepth: 9\n"  >> $(TARGET_DIR_PRODUCER)/index.rst
	$(PRINTF) "\n"                 >> $(TARGET_DIR_PRODUCER)/index.rst

prep_consumers:
	mkdir -p $(TARGET_DIR_CONSUMER)
	if [ -e $(TARGET_DIR_CONSUMER)/index.rst ]; then rm $(TARGET_DIR_CONSUMER)/index.rst ; fi

	$(PRINTF) ".. Autogenerated by docs/Makefile\n"         >> $(TARGET_DIR_CONSUMER)/index.rst
	$(PRINTF) ".. include:: ../../plugins/consumer.rst\n" >> $(TARGET_DIR_CONSUMER)/index.rst
	$(PRINTF) "\n"                 >> $(TARGET_DIR_CONSUMER)/index.rst
	$(PRINTF) ".. toctree::\n"     >> $(TARGET_DIR_CONSUMER)/index.rst
	$(PRINTF) "    :maxdepth: 9\n"  >> $(TARGET_DIR_CONSUMER)/index.rst
	$(PRINTF) "\n"                 >> $(TARGET_DIR_CONSUMER)/index.rst

prep_filters:
	mkdir -p $(TARGET_DIR_FILTER)
	if [ -e $(TARGET_DIR_FILTER)/index.rst ]; then rm $(TARGET_DIR_FILTER)/index.rst ; fi

	$(PRINTF) ".. Autogenerated by docs/Makefile\n"         >> $(TARGET_DIR_FILTER)/index.rst
	$(PRINTF) ".. include:: ../../plugins/filter.rst\n" >> $(TARGET_DIR_FILTER)/index.rst
	$(PRINTF) "\n"                 >> $(TARGET_DIR_FILTER)/index.rst
	$(PRINTF) ".. toctree::\n"     >> $(TARGET_DIR_FILTER)/index.rst
	$(PRINTF) "    :maxdepth: 9\n"  >> $(TARGET_DIR_FILTER)/index.rst
	$(PRINTF) "\n"                 >> $(TARGET_DIR_FILTER)/index.rst

prep_formatters:
	mkdir -p $(TARGET_DIR_FORMATTER)
	if [ -e $(TARGET_DIR_FORMATTER)/index.rst ]; then rm $(TARGET_DIR_FORMATTER)/index.rst ; fi

	$(PRINTF) ".. Autogenerated by docs/Makefile\n"         >> $(TARGET_DIR_FORMATTER)/index.rst
	$(PRINTF) ".. include:: ../../plugins/formatter.rst\n" >> $(TARGET_DIR_FORMATTER)/index.rst
	$(PRINTF) "\n"                 >> $(TARGET_DIR_FORMATTER)/index.rst
	$(PRINTF) ".. toctree::\n"     >> $(TARGET_DIR_FORMATTER)/index.rst
	$(PRINTF) "    :maxdepth: 9\n"  >> $(TARGET_DIR_FORMATTER)/index.rst
	$(PRINTF) "\n"                 >> $(TARGET_DIR_FORMATTER)/index.rst

prep_routers:
	mkdir -p $(TARGET_DIR_ROUTER)
	if [ -e $(TARGET_DIR_ROUTER)/index.rst ]; then rm $(TARGET_DIR_ROUTER)/index.rst ; fi

	$(PRINTF) ".. Autogenerated by docs/Makefile\n"         >> $(TARGET_DIR_ROUTER)/index.rst
	$(PRINTF) ".. include:: ../../plugins/router.rst\n" >> $(TARGET_DIR_ROUTER)/index.rst
	$(PRINTF) "\n"                 >> $(TARGET_DIR_ROUTER)/index.rst
	$(PRINTF) ".. toctree::\n"     >> $(TARGET_DIR_ROUTER)/index.rst
	$(PRINTF) "    :maxdepth: 9\n"  >> $(TARGET_DIR_ROUTER)/index.rst
	$(PRINTF) "\n"                 >> $(TARGET_DIR_ROUTER)/index.rst

# Loop targets for the .rst files
$(rsts_producer):  $(RST_GENERATOR)
	$(RST_GENERATOR) $(MODULES_DIR_PRODUCER)/$(shell basename $@ .rst).go $@
	$(PRINTF) "    %s\n" /$(TARGET_DIR_PRODUCER)/$(shell basename $@) >> $(TARGET_DIR_PRODUCER)/index.rst

$(rsts_consumer):  $(RST_GENERATOR)
	$(RST_GENERATOR) $(MODULES_DIR_CONSUMER)/$(shell basename $@ .rst).go $@
	$(PRINTF) "    %s\n" /$(TARGET_DIR_CONSUMER)/$(shell basename $@) >> $(TARGET_DIR_CONSUMER)/index.rst

$(rsts_filter):  $(RST_GENERATOR)
	$(RST_GENERATOR) $(MODULES_DIR_FILTER)/$(shell basename $@ .rst).go $@
	$(PRINTF) "    %s\n" /$(TARGET_DIR_FILTER)/$(shell basename $@) >> $(TARGET_DIR_FILTER)/index.rst

$(rsts_formatter):  $(RST_GENERATOR)
	$(RST_GENERATOR) $(MODULES_DIR_FORMATTER)/$(shell basename $@ .rst).go $@
	$(PRINTF) "    %s\n" /$(TARGET_DIR_FORMATTER)/$(shell basename $@) >> $(TARGET_DIR_FORMATTER)/index.rst

$(rsts_router):  $(RST_GENERATOR)
	$(RST_GENERATOR) $(MODULES_DIR_ROUTER)/$(shell basename $@ .rst).go $@
	$(PRINTF) "    %s\n" /$(TARGET_DIR_ROUTER)/$(shell basename $@) >> $(TARGET_DIR_ROUTER)/index.rst

####################################################################################################

#
# Sphinx target for generating HTML docs from the .rst files
#
# This is only needed for testing sphinx generation locally, e.g. to verify .rst syntax. Actual
# documentation is generated from the .rst files in the github repository automatically by
# readthedocs.org.
#
# To install Sphinx locally, do:
#   pip install -U Sphinx
#   pip install sphinx_rtd_theme
#

HTML_BUILD_DIR=_build.html

$(HTML_BUILD_DIR):
	mkdir $(HTML_BUILD_DIR)
	
clean-html:
	if [ -e $(HTML_BUILD_DIR) ]; then rm -r $(HTML_BUILD_DIR) ; fi

html: $(HTML_BUILD_DIR)
	mkdir -p $(HTML_BUILD_DIR)/doctrees $(HTML_BUILD_DIR)/html
	sphinx-build -b html -d $(HTML_BUILD_DIR)/doctrees . $(HTML_BUILD_DIR)/html

####################################################################################################
