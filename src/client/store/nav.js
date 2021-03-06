export const state = () => ({
  drawerState: false,
})

export const mutations = {
  toggleDrawer(state) {
    state.drawerState = !state.drawerState
  },
}
