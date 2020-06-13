import axios from 'axios'
import { FETCH_RESOURCE, DELETE_RESOURCE } from "../types";

export const fetchResource = () => {
  return async (dispatch) => {
    const response = await axios.get("http://192.168.0.16:8080/api/v1/resources");
    dispatch({ type: FETCH_RESOURCE, payload: response.data });
  };
};

export const deleteResource = (context) => {
  const { name } = context
  return async (dispatch) => {
    const response = await axios.delete(`http://192.168.0.16:8080/api/v1/resources/${name}`);
    dispatch({ type: DELETE_RESOURCE, payload: response.data });
  };
};