import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/views/Index.vue'
import Types from '@/views/Types.vue'
import NFT from '@/views/NFT.vue'
import TrustAnchors from '@/views/TrustAnchors.vue'
import Relayers from '@/views/Relayers.vue'

const routerHistory = createWebHistory()
const routes = [
  {
    path: '/',
    component: Index,
  },
  { path: '/types', component: Types },
  { path: '/relayers', component: Relayers },
  { path: '/nfts', component: NFT },
  { path: '/trust-anchors', component: TrustAnchors },
]

const router = createRouter({
  history: routerHistory,
  routes,
})

export default router
