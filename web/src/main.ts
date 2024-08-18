import { createApp } from "vue";
import App from "./App.vue"
import "./assets/tailwind.css";
import './assets/common.less'
import {AppName, StroageCache} from '@/utils/config';

import {init as initStorage} from '@/libs/localCache.js'

async function init() {
  await initStorage(AppName, StroageCache)
  const app = createApp(App);
  app.mount("#app");
}
init()
