import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "../../../configs/AxiosConfig";

interface PersonalData {
  nickname: string;
  oneLineIntroduce: string;
  hashtag: string[];
}

const initialState: PersonalData = {
  nickname: "",
  oneLineIntroduce: "",
  hashtag: [],
};

export const getPersonalData = createAsyncThunk<
  PersonalData,
  string | undefined
>("personal_data", async (id) => {
  const response = await axios.get<PersonalData>(
    `/users/profiles/${id}/personaldatas`,
    {
      withCredentials: true,
    }
  );
  return response.data;
});

const personalDateSlice = createSlice({
  name: "personal_data",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(getPersonalData.fulfilled, (state, action) => {
        state.nickname = action.payload.nickname;
        state.oneLineIntroduce = action.payload.oneLineIntroduce;
        if (action.payload?.hashtag == undefined) {
          state.hashtag = [];
        } else {
          state.hashtag = action.payload?.hashtag;
        }
      })
      .addCase(getPersonalData.rejected, (state, action) => {
        console.error("Error fetching user info:", action.error.message);
      });
  },
});
export default personalDateSlice.reducer;
