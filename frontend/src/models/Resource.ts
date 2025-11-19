import mongoose, { Schema } from "mongoose";

const AccessRequestSchema = new Schema(
  {
    user: { type: Schema.Types.ObjectId, ref: "User", required: true },
    message: { type: String, default: "" },
    status: { type: String, enum: ["pending", "approved", "rejected"], default: "pending" },
    requestedAt: { type: Date, default: Date.now },
    respondedAt: { type: Date, default: null },
  },
  { _id: true }
);

const ResourceSchema = new Schema(
  {
    owner: { type: Schema.Types.ObjectId, ref: "User", required: true },

    title: { type: String, required: true },
    description: { type: String, default: "" },
    category: { type: String, default: "Other" },
    subject: { type: String, default: "" },
    tags: { type: [String], default: [] },

    fileName: { type: String, required: true },
    fileUrl: { type: String, required: true },
    
    isPublic: { type: Boolean, default: true },

    approvedUsers: [{ type: Schema.Types.ObjectId, ref: "User" }],
    requests: [AccessRequestSchema],
  },
  { timestamps: true }
);

export default mongoose.models.Resource || mongoose.model("Resource", ResourceSchema);
