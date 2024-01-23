import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "../../../configs/AxiosConfig";
import { TechData } from "../../../GlobalType/Tech";

interface TechList {
  tech_list: TechData[];
}
const initialState: TechList = {
  tech_list: [],
};

export const getTechList = createAsyncThunk<TechList, string | undefined>(
  "tech_data/getTechList",
  async (id) => {
    const response = await axios.get<TechList>(`/profile/techs/${id}`, {
      withCredentials: true,
    });
    return response.data;
  }
);

export const setTechList = createAsyncThunk<TechList>(
  "tech_data/setTechList",
  async (id, tech_list) => {
    const techListJsonString = JSON.stringify(tech_list);

    const response = await axios.post<TechList>(
      `/profile/techs/${id}`,
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
      .addCase(getTechList.fulfilled, (state, action) => {
        if (
          action.payload?.tech_list == undefined ||
          action.payload?.tech_list == null
        ) {
          state.tech_list = [];
        } else {
          state.tech_list = action.payload?.tech_list;
        }
      })
      .addCase(getTechList.rejected, (state, action) => {
        console.error("Error get tech list:", action.error.message);
      })
      .addCase(setTechList.fulfilled, (state, action) => {
        if (action.payload.tech_list == undefined) {
          state.tech_list = [];
        } else {
          state.tech_list = action.payload.tech_list;
        }
      })
      .addCase(setTechList.rejected, (state, action) => {
        console.error("Error get tech list:", action.error.message);
      });
  },
});
export default techDataSlice.reducer;
