<template>
  <div></div>
</template>

<script>
import Loggly from "loggly-jslogger";

export default {
  name: "LogglyLogger",
  created() {
    const logger = Loggly.createLogger({
      logglyKey: "e5efab8f-e732-4bba-9a1f-ad572a6c3a47",
      sendConsoleErrors: true,
      tag: "FrontEndLogs",
    });
    logger.log("Hello from Vue.js");
  },
};
</script>
