import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "../../../configs/AxiosConfig";
import { TechData } from "../../../GlobalType/Tech";

interface TechList {
  tech_list: TechData[];
}
const initialState: TechList = {
  tech_list: [],
};

export const getUserTechList = createAsyncThunk<TechList, string | undefined>(
  "tech_data/getUserTechList",
  async (id) => {
    const response = await axios.get<TechList>(`/users/profiles/${id}/techs`, {
      withCredentials: true,
    });
    console.log(response.data);
    return response.data;
  }
);

export const setUserTechList = createAsyncThunk<TechList>(
  "tech_data/setUserTechList",
  async (id, tech_list) => {
    const techListJsonString = JSON.stringify(tech_list);

    const response = await axios.put<TechList>(
      `/users/profiles/${id}/techs`,
      techListJsonString,
      {
        withCredentials: true,
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    return response.data;
  }
);

const techDataSlice = createSlice({
  name: "tech_data",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(getUserTechList.fulfilled, (state, action) => {
        console.log(action.payload);
        if (action.payload == undefined || action.payload == null) {
          state.tech_list = [];
        } else {
          state.tech_list = action.payload.tech_list;
        }
      })
      .addCase(getUserTechList.rejected, (state, action) => {
        console.error("Error get tech list:", action.error.message);
      })
      .addCase(setUserTechList.fulfilled, (state, action) => {
        if (action.payload.tech_list == undefined) {
          state.tech_list = [];
        } else {
          state.tech_list = action.payload.tech_list;
        }
      })
      .addCase(setUserTechList.rejected, (state, action) => {
        console.error("Error get tech list:", action.error.message);
      });
  },
});
export default techDataSlice.reducer;
