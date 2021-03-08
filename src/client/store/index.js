export const actions = {
  async nuxtServerInit({ dispatch }) {
    await dispatch('company/loadInfo')
  },
}
