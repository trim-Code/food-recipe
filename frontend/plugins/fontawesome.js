import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faSearch,
  faBookmark,
  faUser,
  faBars,
  faTimes,
  faArrowDown,
  faShoppingCart,
  faPlus,
} from "@fortawesome/free-solid-svg-icons";

library.add(
  faSearch,
  faBookmark,
  faUser,
  faBars,
  faTimes,
  faArrowDown,
  faShoppingCart,
  faPlus
);

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.component("font-awesome-icon", FontAwesomeIcon);
});
