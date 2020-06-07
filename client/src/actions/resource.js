import axios from 'axios'
import { FETCH_RESOURCE } from "../types";

export const resourceAction = () => {
  return async (dispatch) => {
    const response = await axios.get("http://192.168.0.16:8080/api/resources");
    dispatch({ type: FETCH_RESOURCE, payload: response.data });
  };
};
