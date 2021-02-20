import Layout from '@/layout'

const nestedRouter = {
  path: '/kubernetes',
  component: Layout,
  redirect: '/kubernetes/latest',
  name: 'Nested',
  meta: {
    title: 'Kubernetes',
    icon: 'nested'
  },
  children: [
    {
      path: 'latest',
      name: 'latest',
      component: () => import('@/views/kubernetes/latest/index'),
      meta: { title: 'latest' }
    },
    {
      path: 'stage',
      name: 'stage',
      component: () => import('@/views/nested/menu2/index'),
      meta: { title: 'stage' }
    },
    {
      path: 'lt01',
      name: 'lt01',
      component: () => import('@/views/nested/menu2/index'),
      meta: { title: 'lt01' }
    },
    {
      path: 'prod',
      name: 'prod',
      component: () => import('@/views/nested/menu2/index'),
      meta: { title: 'prod' }
    },
    {
      path: 'proda',
      name: 'proda',
      component: () => import('@/views/nested/menu2/index'),
      meta: { title: 'proda' }
    },
    {
      path: 'prodb',
      name: 'prodb',
      component: () => import('@/views/nested/menu2/index'),
      meta: { title: 'prodb' }
    }
  ]
}

export default nestedRouter
