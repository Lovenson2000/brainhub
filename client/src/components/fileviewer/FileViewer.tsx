"use client";

import * as React from "react";
import { Viewer, Worker } from "@react-pdf-viewer/core";
import { defaultLayoutPlugin } from "@react-pdf-viewer/default-layout";
import "@react-pdf-viewer/core/lib/styles/index.css";
import "@react-pdf-viewer/default-layout/lib/styles/index.css";

export default function FileViewer() {
  const defaultLayoutPluginInstance = defaultLayoutPlugin();

  const filePath = "/assets/Lovenson_Beaucicot_Resume.pdf";

  return (
    <Worker workerUrl="https://unpkg.com/pdfjs-dist@3.11.174/build/pdf.worker.min.js">
      <div className="mx-auto h-[45rem] w-[40rem]">
        <Viewer fileUrl={filePath} plugins={[defaultLayoutPluginInstance]} />
      </div>
    </Worker>
  );
}
