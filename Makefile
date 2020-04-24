include build/makefiles/pkg/base/base.mk
include build/makefiles/pkg/string/string.mk
include build/makefiles/pkg/color/color.mk
include build/makefiles/pkg/functions/functions.mk
include build/makefiles/target/buildenv/buildenv.mk
include build/makefiles/target/go/go.mk
# include build/makefiles/target/tests/header/header.mk
# include build/makefiles/target/tests/config/config.mk
# include build/makefiles/target/tests/overlay-network/overlay-network.mk
THIS_FILE := $(firstword $(MAKEFILE_LIST))
SELF_DIR := $(dir $(THIS_FILE))
.PHONY: test build clean run kill proto config  
.SILENT: test build clean run kill proto config 
CONFIG_DIR:=$(PWD)/fixtures
PORT_ONE:=8080 
PORT_TWO:=8081
PORT_THREE:=8082
PORT_FOUR:=8083
build: 
	- $(call print_running_target)
	- @$(MAKE) --no-print-directory -f $(THIS_FILE) go-build
	- $(call print_completed_target)
proto: 
	- $(call print_running_target)
	- @$(MAKE) --no-print-directory -C model proto
	- $(call print_completed_target)
run: kill
	- $(call print_running_target)
	- $(MKDIR) logs && bin$(PSEP)overlay-network daemon --dev --log-level=debug -config-file=$(PWD)/fixtures/rc1.json --rpc-port=${PORT_TWO} > $(PWD)/logs/rc1.log 2>&1 &
	- $(call print_completed_target)
config: 
	- $(call print_running_target)
	- $(info $(CONFIG_DIR))
	- bin$(PSEP)overlay-network parse-config -config-dir=$(CONFIG_DIR)
	- $(call print_completed_target)
clean: 
	- $(call print_running_target)
	- @$(MAKE) --no-print-directory -f $(THIS_FILE) go-clean
	- $(call print_completed_target)
kill : 
	- $(call print_running_target)
	- $(RM) $(PWD)/server.log
	- for pid in $(shell ps  | grep "overlay-network" | awk '{print $$1}'); do kill -9 "$$pid"; done
	- $(call print_completed_target)
