import { createStore, combineReducers, applyMiddleware } from "redux";
import ReduxThunk from "redux-thunk";
import { resourceReducer } from '../reducers'

const rootReducer = combineReducers({
  resources: resourceReducer
})

export default createStore(rootReducer, applyMiddleware(ReduxThunk));