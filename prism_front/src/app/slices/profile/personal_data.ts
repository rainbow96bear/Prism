import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "../../../configs/AxiosConfig";

interface PersonalData {
  nickname: string;
  profile_img: string;
  one_line_introduce: string;
  hashtag: string[];
}

const initialState: PersonalData = {
  nickname: "",
  profile_img:
    process.env.REACT_APP_BASE_URL + "/assets/base_profile.jpeg" || "",
  one_line_introduce: "한 줄 소개",
  hashtag: ["golang"],
};

export const getPersonalDate = createAsyncThunk<
  PersonalData,
  string | undefined
>("personal_data", async (id) => {
  const response = await axios.get<PersonalData>(
    `/profile/personaldata/${id}`,
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
      .addCase(getPersonalDate.fulfilled, (state, action) => {
        state.nickname = action.payload.nickname;
        state.profile_img = action.payload.profile_img;
        state.one_line_introduce = action.payload.one_line_introduce;
        // state.hashtag = action.payload.hashtag;
      })
      .addCase(getPersonalDate.rejected, (state, action) => {
        console.error("Error fetching user info:", action.error.message);
      });
  },
});
export default personalDateSlice.reducer;
