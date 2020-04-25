include build/makefiles/pkg/base/base.mk
include build/makefiles/pkg/string/string.mk
include build/makefiles/pkg/color/color.mk
include build/makefiles/pkg/functions/functions.mk
include build/makefiles/target/buildenv/buildenv.mk
include build/makefiles/target/go/go.mk
THIS_FILE := $(firstword $(MAKEFILE_LIST))
SELF_DIR := $(dir $(THIS_FILE))
.PHONY: test build clean run kill proto config  
.SILENT: test build clean run kill proto config 
CONFIG_DIR:=$(PWD)/fixtures
PORT_ONE:=8080 
PORT_TWO:=8081
PORT_THREE:=8082
DELAY=1
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
	- @$(MAKE) --no-print-directory -f $(THIS_FILE) run-1
	- sleep $(DELAY)
	- @$(MAKE) --no-print-directory -f $(THIS_FILE) run-2
	- $(eval DELAY=$(shell echo $$(($(DELAY)+1))))
	- sleep $(DELAY)
	- @$(MAKE) --no-print-directory -f $(THIS_FILE) run-3
	- $(eval DELAY=$(shell echo $$(($(DELAY)+1))))
	- sleep $(DELAY)
	- $(call print_completed_target)


.PHONY:  run-3 run-2 run-1
.SILENT: run-3 run-2 run-1

run-3: 
	- $(call print_running_target)
	- $(MKDIR) logs && bin$(PSEP)overlay-network daemon -config-file=$(PWD)/fixtures/rc1.json --rpc-port=${PORT_THREE} > $(PWD)/logs/node-3.log 2>&1 &
	- $(call print_completed_target)
run-2: 
	- $(call print_running_target)
	- $(MKDIR) logs && bin$(PSEP)overlay-network daemon -config-file=$(PWD)/fixtures/rc1.json --rpc-port=${PORT_TWO} > $(PWD)/logs/node-2.log 2>&1 &
	- $(call print_completed_target)
run-1: 
	- $(call print_running_target)
	- $(MKDIR) logs && bin$(PSEP)overlay-network daemon -config-file=$(PWD)/fixtures/rc1.json --rpc-port=${PORT_ONE} > $(PWD)/logs/node-1.log 2>&1 &
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
