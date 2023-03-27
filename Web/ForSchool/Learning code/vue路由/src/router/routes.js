import A from "./../components/A.vue";
import B from "./../components/B.vue";
import C from "./../components/C.vue";
import D from "./../components/D.vue";

const routes = [
  {
    path: "/a",
    name: "A",
    component: A,
    children: [
      {
        path: "c",
        component: C,
      },
      {
        path: "d",
        component: D,
      }
    ]
  },
  {
    path: "/b",
    name: "B",
    component: B,
  },
];

export default routes;