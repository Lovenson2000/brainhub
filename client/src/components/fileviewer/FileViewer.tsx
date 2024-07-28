"use client";

import * as React from "react";
import { Viewer, Worker, Button } from "@react-pdf-viewer/core";
import { defaultLayoutPlugin } from "@react-pdf-viewer/default-layout";
import { highlightPlugin, Trigger, RenderHighlightTargetProps } from "@react-pdf-viewer/highlight";
import "@react-pdf-viewer/core/lib/styles/index.css";
import "@react-pdf-viewer/default-layout/lib/styles/index.css";
import "@react-pdf-viewer/highlight/lib/styles/index.css";

interface HighlightArea {
  height: number;
  left: number;
  pageIndex: number;
  top: number;
  width: number;
}

interface Note {
  id: number;
  content: string;
  highlightAreas: HighlightArea[];
  quote: string;
}

export default function FileViewer() {
  const defaultLayoutPluginInstance = defaultLayoutPlugin();

  const [message, setMessage] = React.useState("");
  const [notes, setNotes] = React.useState<Note[]>([]);
  let noteId = notes.length;

  const renderHighlightTarget = (props: RenderHighlightTargetProps) => {
    console.log("Selected text:", props.selectedText);

    const addNote = () => {
      if (message !== "") {
        const note: Note = {
          id: ++noteId,
          content: message,
          highlightAreas: props.highlightAreas,
          quote: props.selectedText,
        };
        setNotes(notes.concat([note]));

        props.cancel();
      }
    };

    return (
      <div
        style={{
          background: "#fff",
          border: "1px solid rgba(0, 0, 0, .3)",
          borderRadius: "2px",
          padding: "8px",
          position: "absolute",
          left: `${props.selectionRegion.left}%`,
          top: `${props.selectionRegion.top + props.selectionRegion.height}%`,
          zIndex: 1,
        }}
      >
        <div>
          <textarea
            rows={3}
            style={{
              border: "1px solid rgba(0, 0, 0, .3)",
            }}
            onChange={(e) => setMessage(e.target.value)}
          ></textarea>
        </div>
        <div
          style={{
            display: "flex",
            marginTop: "8px",
          }}
        >
          <div style={{ marginRight: "8px" }}>
            <Button onClick={addNote}>Add</Button>
          </div>
          <Button onClick={props.cancel}>Cancel</Button>
        </div>
      </div>
    );
  };

  const highlightPluginInstance = highlightPlugin({
    renderHighlightTarget,
    trigger: Trigger.TextSelection,
  });

  const filePath =
    "/assets/James Stewart, Daniel K. Clegg, Saleem Watson, Lothar Redlin - Calculus_ Early Transcendentals. 9e-Cengage Learning (2020).pdf";

  return (
    <Worker workerUrl="https://unpkg.com/pdfjs-dist@3.11.174/build/pdf.worker.min.js">
      <div className="mx-auto h-[45rem] w-[40rem]">
        <Viewer
          fileUrl={filePath}
          plugins={[defaultLayoutPluginInstance, highlightPluginInstance]}
        />
      </div>
    </Worker>
  );
}
