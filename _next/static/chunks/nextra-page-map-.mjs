import meta from "../../../src/pages/_meta.ts";
import development_meta from "../../../src/pages/development/_meta.ts";
import getting_started_meta from "../../../src/pages/getting_started/_meta.ts";
export const pageMap = [{
  data: meta
}, {
  name: "development",
  route: "/development",
  children: [{
    data: development_meta
  }, {
    name: "index",
    route: "/development",
    frontMatter: {
      "title": "Quickstart"
    }
  }]
}, {
  name: "getting_started",
  route: "/getting_started",
  children: [{
    data: getting_started_meta
  }, {
    name: "concepts",
    route: "/getting_started/concepts",
    frontMatter: {
      "title": "Concepts"
    }
  }, {
    name: "index",
    route: "/getting_started",
    frontMatter: {
      "title": "Quickstart"
    }
  }]
}, {
  name: "index",
  route: "/",
  frontMatter: {
    "title": "Home"
  }
}];