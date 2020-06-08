import axios from 'axios'
import { FETCH_RESOURCE, DELETE_RESOURCE } from "../types";

export const fetchResource = () => {
  return async (dispatch) => {
    const response = await axios.get("http://192.168.0.16:8080/api/resources");
    dispatch({ type: FETCH_RESOURCE, payload: response.data });
  };
};

export const deleteResource = ({ selfLink }) => {
  return async (dispatch) => {
    const response = await axios.post(`http://192.168.0.16:8080/api/resources?selfLink=${selfLink}`);
    dispatch({ type: DELETE_RESOURCE, payload: response.data });
  };
};