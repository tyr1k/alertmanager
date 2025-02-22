{-
   Alertmanager API
   API of the Prometheus Alertmanager (https://github.com/tyr1k/alertmanager)

   OpenAPI spec version: 0.0.1

   NOTE: This file is auto generated by the openapi-generator.
   https://github.com/openapitools/openapi-generator.git
   Do not edit this file manually.
-}


module Data.Matcher exposing (Matcher, decoder, encoder)

import Dict exposing (Dict)
import Json.Decode as Decode exposing (Decoder)
import Json.Decode.Pipeline exposing (optional, required)
import Json.Encode as Encode


type alias Matcher =
    { name : String
    , value : String
    , isRegex : Bool
    , isEqual : Maybe Bool
    }


decoder : Decoder Matcher
decoder =
    Decode.succeed Matcher
        |> required "name" Decode.string
        |> required "value" Decode.string
        |> required "isRegex" Decode.bool
        |> optional "isEqual" (Decode.nullable Decode.bool) (Just True)


encoder : Matcher -> Encode.Value
encoder model =
    Encode.object
        [ ( "name", Encode.string model.name )
        , ( "value", Encode.string model.value )
        , ( "isRegex", Encode.bool model.isRegex )
        , ( "isEqual", Maybe.withDefault Encode.null (Maybe.map Encode.bool model.isEqual) )
        ]
