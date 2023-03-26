import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    // children: [
    //   {
    //     path: "/edu/student",
    //     name: "Student",
    //     component: () => import('../views/Student.vue')
    //   },
    //   {
    //     path: "/edu/class",
    //     name: "Class",
    //     component: () => import('../views/Class.vue')
    //   },
    //   {
    //     path: "/edu/course",
    //     name: "Course",
    //     component: () => import('../views/Course.vue')
    //   }
    // ]
    children: [
      {
        path: "/edu/student",
        name: "Student",
        component: () => import('../views/javaweb/Student.vue')
      },
      {
        path: "/edu/class",
        name: "Class",
        component: () => import('../views/javaweb/Class.vue')
      },
      {
        path: "/edu/course",
        name: "Course",
        component: () => import('../views/javaweb/Course.vue')
      }
    ]
  },
  {
    path: '/about',
    name: 'About',
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
