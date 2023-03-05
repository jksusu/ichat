import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
import HomeView from '@/renderer/views/home/Index.vue'
import LoginView from '@/renderer/views/login/Login.vue'
import sessionView from '@/renderer/views/home/Session.vue'
import featureHome from '@/renderer/views/friends/Index.vue'
import settingView from '@/renderer/views/setting/Index.vue'
import MessageTrip from '@/renderer/views/other/MessageTrip.vue'
import NewFriendsView from '@/renderer/views/friends/NewFriends.vue'
import GroupListsView from '@/renderer/views/friends/GroupLists.vue'
import FriendListsView from '@/renderer/views/friends/FriendLists.vue'
import ScanCodeLoginView from '@/renderer/views/login/ScanCodeLogin.vue'
import RetrievePasswordView from '@/renderer/views/login/RetrievePassword.vue'


const routes: Array<RouteRecordRaw> = [
  { path: '/', component: sessionView, meta: { auth: true, keepAlive: true } },
  { path: '/home', component: HomeView, meta: { auth: true } },
  { path: '/login', component: LoginView, meta: { auth: false } },
  { path: '/scan_code_login', component: ScanCodeLoginView },
  { path: '/retrieve_password', component: RetrievePasswordView },
  { path: '/register', component: LoginView },
  { path: '/session', component: sessionView, meta: { auth: true, keepAlive: true } },
  { path: '/messageTrip', component: MessageTrip },
  {
    path: '/maillist', component: featureHome, meta: { auth: true },
    children: [
      // { path: 'maillist', component: featureHome, meta: { auth: true } },
      { path: 'new_friends', component: NewFriendsView, meta: { auth: true } },
      { path: 'groups', component: GroupListsView, meta: { auth: true } },
      { path: 'friends', component: FriendListsView, meta: { auth: true } },
    ]
  },
  { path: '/setting', component: settingView, meta: { auth: true } },
]

export const router = createRouter({
  history: createWebHashHistory(),
  routes
})