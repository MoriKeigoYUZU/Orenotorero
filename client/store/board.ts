import { Module, MutationTree, GetterTree, ActionTree } from 'vuex'
import { RootState } from '~/store/index'

export interface BoardState {
  boardData: Array<any>
}
export const state: () => BoardState = (): BoardState => ({
  boardData: [
    {
      id: 1,
      title: '今週やること',
      position: 1,
      card: [
        {
          id: 1,
          title:
            '牛乳を買いに行くんじゃあああああああああああああああああああああああああああああああああああああ',
          describe: 'これは1番目のカードです。',
          deadline: '2019-05-03 11:00',
          position: 1
        },
        {
          id: 2,
          title: '牛を狩に行く',
          describe: 'これは1番目のカードです。',
          deadline: '2019-05-03 11:00',
          position: 2
        }
      ]
    },
    {
      id: 2,
      title: 'TODO',
      position: 2,
      card: [
        {
          id: 5,
          title: 'ミルクゲット',
          describe: 'これは1番目のカードです。',
          deadline: '2019-05-03 11:00',
          position: 1
        },
        {
          id: 3,
          title: 'ミートゲット',
          describe: 'これは1番目のカードです。',
          deadline: '2019-05-03 11:00',
          position: 2
        }
      ]
    }
  ]
})
const getters: GetterTree<BoardState, RootState> = {
  boardData(state: BoardState) {
    return state.boardData
  },
  kanbanData: (state: BoardState) => (index: number) => {
    return state.boardData[index].card
  }
}
const mutations: MutationTree<BoardState> = {
  updateBoardData(state: BoardState, newBoardData: any): void {
    state.boardData = newBoardData
  },
  updateKanbanData(state: BoardState, payload: any) {
    state.boardData[payload.index].card = payload.value
  }
}
const actions: ActionTree<BoardState, RootState> = {
  async getBoardData({ commit }: any, payload: object): Promise<any> {
    console.log(payload)
    await this.$axios
      .get('/kanban', {
        headers: {
          Authorization: 'Bearer ' + payload.token,
          BoardId: payload.boardId
        }
      })
      .then((res: any) => {
        console.log(res.data)
        commit('updateBoardData', res.data)
      })
      .catch((err: any) => {
        console.log(err)
      })
  },
  async updateCardTitle({ commit }: any, payload: object): Promise<any> {
    console.log('payload:', payload)
    await this.$axios
      .put(
        '/card',
        {
          id: payload.id,
          title: payload.title
        },
        {
          headers: {
            Authorization: 'Bearer ' + payload.token
          }
        }
      )
      .catch((err: any) => {
        console.log(err)
      })
  }
}
export const Board: Module<BoardState, RootState> = {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
export default Board
