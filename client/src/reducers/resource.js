import { FETCH_RESOURCE, DELETE_RESOURCE } from "../types";

const initialState = {
  resources: [],
};

export const resourceReducer = (state = initialState, action) => {
  switch (action.type) {
    case FETCH_RESOURCE:
      return {
        ...state,
        resources: action.payload.data,
      };
    case DELETE_RESOURCE:
      return {
        ...state,
        resources: action.payload.data,
      };
    default:
      return state;
  }
};
