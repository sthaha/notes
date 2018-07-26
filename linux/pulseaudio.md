## Fix bluetooth headphone not getting added to device list

References:
  * https://bugzilla.freedesktop.org/show_bug.cgi?id=73325#c52
  * https://bugzilla.redhat.com/show_bug.cgi?id=1503759

```diff
diff --git a/pulse/default.pa b/pulse/default.pa
index bbcbee8..79ab05e 100644
--- a/etc/pulse/default.pa
+++ b/etc/pulse/default.pa
@@ -61,7 +61,7 @@ load-module module-bluetooth-policy
 .endif

 .ifexists module-bluetooth-discover.so
-load-module module-bluetooth-discover
+load-module module-bluetooth-discover headset=auto
 .endif

```

