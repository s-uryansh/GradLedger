import mongoose, { Schema, Document } from 'mongoose';

export interface IOtp extends Document {
  email: string;
  code: string;
  createdAt: Date;
}

const OtpSchema = new Schema<IOtp>({
  email: { type: String, required: true },
  code: { type: String, required: true },
  createdAt: { type: Date, default: Date.now, expires: 300 } 
});

export default mongoose.models.Otp || mongoose.model<IOtp>('Otp', OtpSchema);
