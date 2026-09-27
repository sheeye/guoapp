package main

/*
#include <signal.h>
#include <execinfo.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <time.h>

#if defined(__ANDROID__) || defined(__linux__)
static char crash_dir[1024];
static int have_dir = 0;

static const char* signal_name(int sig) {
    switch (sig) {
        case SIGILL: return "SIGILL";
        case SIGSEGV: return "SIGSEGV";
        case SIGBUS: return "SIGBUS";
        case SIGABRT: return "SIGABRT";
        case SIGTRAP: return "SIGTRAP";
        case SIGFPE: return "SIGFPE";
        case SIGSYS: return "SIGSYS";
        default: return "UNKNOWN";
    }
}

// 通过 /proc/self/maps 把地址解析到所属动态库，便于判断是 Go 核心还是播放器。
static void resolve_lib(void* addr, char* out, size_t outlen) {
    FILE* maps = fopen("/proc/self/maps", "r");
    if (!maps) { snprintf(out, outlen, "?"); return; }
    char line[512];
    unsigned long target = (unsigned long)addr;
    while (fgets(line, sizeof(line), maps)) {
        unsigned long start, end, offset;
        char perms[8];
        char path[256];
        if (sscanf(line, "%lx-%lx %7s %lx %*x:%*x %*d %255[^\n]",
                   &start, &end, perms, &offset, path) == 5) {
            if (target >= start && target < end) {
                snprintf(out, outlen, "%s (offset 0x%lx)", path, target - start);
                fclose(maps);
                return;
            }
        }
    }
    fclose(maps);
    snprintf(out, outlen, "?");
}

static void crash_handler(int sig, siginfo_t* info, void* ucontext) {
    char path[1100];
    if (have_dir) snprintf(path, sizeof(path), "%s/native_crash.txt", crash_dir);
    else snprintf(path, sizeof(path), "/sdcard/native_crash.txt");
    FILE* f = fopen(path, "w");
    if (f) {
        time_t t = time(NULL);
        fprintf(f, "=== native crash report ===\n");
        fprintf(f, "time=%ld\n", (long)t);
        fprintf(f, "signal=%d (%s)\n", sig, signal_name(sig));
        if (info) {
            fprintf(f, "si_code=%d si_addr=%p\n", info->si_code, info->si_addr);
            char lib[300];
            resolve_lib(info->si_addr, lib, sizeof(lib));
            fprintf(f, "fault_in=%s\n", lib);
        }
        void* bt[64];
        int n = backtrace(bt, 64);
        fprintf(f, "backtrace(%d):\n", n);
        for (int i = 0; i < n; i++) {
            char lib[300];
            resolve_lib(bt[i], lib, sizeof(lib));
            fprintf(f, "  #%d %p in %s\n", i, bt[i], lib);
        }
        fclose(f);
    }
    _exit(1);
}

static void install_crash_handler(void) {
    struct sigaction sa;
    memset(&sa, 0, sizeof(sa));
    sa.sa_sigaction = crash_handler;
    sa.sa_flags = SA_SIGINFO | SA_RESETHAND;
    static char altstack_mem[64 * 1024];
    stack_t st;
    st.ss_sp = altstack_mem;
    st.ss_size = sizeof(altstack_mem);
    sigaltstack(&st, NULL);
    int signals[] = {SIGILL, SIGSEGV, SIGBUS, SIGABRT, SIGTRAP, SIGFPE, SIGSYS, 0};
    for (int i = 0; signals[i]; i++) sigaction(signals[i], &sa, NULL);
}

void duanju_set_crash_dir(const char* dir) {
    if (dir && *dir) {
        strncpy(crash_dir, dir, sizeof(crash_dir) - 1);
        crash_dir[sizeof(crash_dir) - 1] = 0;
        have_dir = 1;
    }
    install_crash_handler();
}

void init_crash_dir_from_env(void) {
    const char* d = getenv("DUANJU_CRASH_DIR");
    if (d && *d) {
        strncpy(crash_dir, d, sizeof(crash_dir) - 1);
        crash_dir[sizeof(crash_dir) - 1] = 0;
        have_dir = 1;
    }
    install_crash_handler();
}
#else
void duanju_set_crash_dir(const char* dir) { (void)dir; }
void init_crash_dir_from_env(void) {}
#endif
*/
import "C"

//export DuanjuSetCrashDir
func DuanjuSetCrashDir(dir *C.char) {
	C.duanju_set_crash_dir(dir)
}

func init() {
	C.init_crash_dir_from_env()
}
