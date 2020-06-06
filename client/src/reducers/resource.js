import { FETCH_RESOURCE } from "../types";

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
    default:
      return state;
  }
};
