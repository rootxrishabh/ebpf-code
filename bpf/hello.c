#include "types.h"
#include "bpf_helpers.h"

SEC("kprobe/sys_clone")
int kprobe_sys_clone(void *ctx) { 
    bpf_printk("Hello, World!\\n"); 
    return 0; 
}

char _license[] SEC("license") = "GPL";
